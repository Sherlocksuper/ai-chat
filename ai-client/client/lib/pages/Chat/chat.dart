import 'package:client/Controller/ChatController.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

import 'chatDetail.dart';
import 'chatItem.dart';

class ChatPage extends StatelessWidget {
  const ChatPage({super.key});

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: Get.find<ChatController>().getChatList(),
      builder: (context, snapshot) {
        return Scaffold(
          appBar: AppBar(
            automaticallyImplyLeading: false,
            title: const Text("Chat with GPT"),
            actions: [
              IconButton(
                  onPressed: () {
                    Get.find<ChatController>().configAI();
                  },
                  icon: const Icon(Icons.add)),
              IconButton(
                  onPressed: () {
                    //弹出弹窗，确认是否删除所有聊天记录
                    Get.defaultDialog(
                      title: 'Delete All Chat',
                      content: const Text('Are you sure to delete all chat?'),
                      actions: [
                        TextButton(
                          onPressed: () {
                            Get.back();
                          },
                          child: const Text('Cancel'),
                        ),
                        TextButton(
                          onPressed: () {
                            Get.find<ChatController>().clearChat();
                            Get.back();
                          },
                          child: const Text('Confirm'),
                        ),
                      ],
                    );
                  },
                  icon: const Icon(Icons.delete_sweep_outlined)),
            ],
          ),
          body: GetBuilder<ChatController>(
            builder: (logic) {
              return RefreshIndicator(
                onRefresh: () async {
                  await Get.find<ChatController>().getChatList();
                },
                child: ListView.separated(
                  itemBuilder: (BuildContext context, int index) {
                    return GestureDetector(
                        onTap: () {
                          Get.to(() => ChatDetail(chat: logic.chatList[index]));
                        },
                        child: ChatItem(chat: logic.chatList[index]));
                  },
                  separatorBuilder: (BuildContext context, int index) {
                    return const Divider();
                  },
                  itemCount: logic.chatList.length,
                ),
              );
            },
          ),
        );
      },
    );
  }
}
