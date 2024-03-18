import 'package:client/model/chat_struct.dart';
import 'package:client/model/message.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:get/get.dart';

import '../Constant.dart';
import '../config.dart';
import 'UserController.dart';

//因为涉及到刷新页面，所以使用GetsController
class ChatController extends GetxController {
  bool isEdited = false;
  List<ChatDetailStruct> chatList = [];

  //配置AI Bot
  void configAI() {
    String title = "默认标题";
    String systemMessage = "你是我的AI助手，我需要你的帮助";
    Get.defaultDialog(
      title: 'AI Config',
      content: Column(
        children: [
          TextField(
            decoration: InputDecoration(hintText: title, border: InputBorder.none),
            onChanged: (value) {
              title = value;
            },
          ),
          TextField(
            decoration: InputDecoration(hintText: systemMessage, border: InputBorder.none),
            onChanged: (value) {
              systemMessage = value;
            },
          ),
        ],
      ),
      actions: [
        TextButton(
          onPressed: () {
            Get.back();
          },
          child: const Text('取消'),
        ),
        TextButton(
          onPressed: () {
            startAChat(title, systemMessage);
            Get.back();
          },
          child: const Text('确定'),
        ),
      ],
    );
  }

  ///开始一个chat
  Future<void> startAChat(String title, String systemMessage) async {
    var map = {'title': title, 'systemMessage': systemMessage, 'userId': UserController.me.id};
    var response = await dio.post(Constant.StartAChatHAT, data: map);
    if (response.data["code"] == 200) {
      getChatList();
    } else {
      EasyLoading.showError(response.data["message"]);
    }
    update();
    print(response);
  }

  ///发送消息
  Future<void> sendMessage(int chatId, String content) async {
    print("chatId:${chatId}content:$content");
    addMessage(chatId, "user", content);

    update();
    var response = await dio.post(Constant.SENDMESSAGE, data: {'chatId': chatId, 'content': content});

    if (response.data["code"] == 200) {
    } else {
      EasyLoading.showError(response.data["message"]);
    }
    update();
  }

  ///获取消息列表
  Future<void> getChatList() async {
    try {
      var response = await dio.get("${Constant.GETCHATLIST}?userId=${UserController.me.id}");
      chatList = (response.data["code"] == 200 || response.data["data"] != null)
          ? List<ChatDetailStruct>.from(response.data["data"].map((e) => ChatDetailStruct.fromJson(e)))
          : throw Exception(response.data["message"]);
      print(chatList.length);
    } catch (e) {
      print(e);
    }
    update();
  }

  ///删除chat
  Future<void> deleteChat(int chatId) async {
    var response = await dio.get("${Constant.DELETECHAT}?id=$chatId");
    if (response.data["code"] == 200) {
      chatList.removeWhere((element) => element.id == chatId);
      update();
    } else {
      EasyLoading.showError(response.data["message"]);
    }
  }

  ///清空chat
  Future<void> clearChat() async {
    var response = await dio.get("${Constant.DELETECHAT}?userId=${UserController.me.id}");
    print(response.data);
    if (response.data["code"] == 200) {
      chatList.clear();
      update();
    } else {
      EasyLoading.showError(response.data["message"]);
    }
  }

  void addMessage(int chatId, String role, String message) {
    print("正在添加数据");
    chatList
        .firstWhere((element) => element.id == chatId)
        .messages
        .add(Message(chatId: chatId, role: role, content: message));
    update([chatId]);
  }
}
