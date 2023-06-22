import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma/prisma.service';
import { InitTransactionDto, InputExecuteTransactionDto } from './order.dto';
import { OrderStatus, OrderType } from '@prisma/client';

@Injectable()
export class OrdersService {
  constructor(private prismaService: PrismaService) {}

  all(filter: { wallet_id: string }) {
    return this.prismaService.order.findMany({
      where: {
        wallet_id: filter.wallet_id,
      },
      include: {
        Transactions: true,
        Asset: {
          select: {
            id: true,
            symbol: true,
          },
        },
      },
      orderBy: {
        updated_at: 'desc',
      },
    });
  }

  initTransaction(input: InitTransactionDto) {
    return this.prismaService.order.create({
      data: {
        wallet_id: input.wallet_id,
        asset_id: input.asset_id,
        shares: input.shares,
        price: input.price,
        type: input.type,
        status: OrderStatus.PENDING,
        partial: input.shares,
      },
    });
  }

  async executeTransaction(input: InputExecuteTransactionDto) {
    return this.prismaService.$transaction(async (prisma) => {
      const order = await prisma.order.findUniqueOrThrow({
        where: { id: input.order_id },
      });

      await prisma.order.update({
        where: { id: input.order_id },
        data: {
          status: input.status,
          partial: order.partial - input.negotiated_shares,
          Transactions: {
            create: {
              broker_transaction_id: input.broker_transaction_id,
              related_investor_id: input.related_investor_id,
              shares: input.negotiated_shares,
              price: input.price,
            },
          },
        },
      });

      if (input.status == OrderStatus.CLOSED) {
        await prisma.asset.update({
          where: { id: order.asset_id },
          data: {
            price: input.price,
          },
        });

        const walletAsset = await prisma.walletAsset.findUnique({
          where: {
            wallet_id_asset_id: {
              asset_id: order.asset_id,
              wallet_id: order.wallet_id,
            },
          },
        });

        if (walletAsset) {
          await prisma.walletAsset.update({
            where: {
              wallet_id_asset_id: {
                asset_id: order.asset_id,
                wallet_id: order.wallet_id,
              },
            },
            data: {
              shares:
                order.type == OrderType.BUY
                  ? walletAsset.shares + input.negotiated_shares
                  : walletAsset.shares - input.negotiated_shares,
            },
          });
        } else {
          await prisma.walletAsset.create({
            data: {
              asset_id: order.asset_id,
              wallet_id: order.wallet_id,
              shares: input.negotiated_shares,
              version: 1,
            },
          });
        }
      }
    });
  }
}
