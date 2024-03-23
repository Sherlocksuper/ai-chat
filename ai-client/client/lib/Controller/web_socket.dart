import 'dart:convert';

import 'package:client/Constant.dart';
import 'package:client/Controller/chat_controller.dart';
import 'package:client/Controller/user_controller.dart';
import 'package:client/model/ws_message.dart';
import 'package:get/get.dart';
import 'package:get/get_core/src/get_main.dart';
import 'package:web_socket_channel/io.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class WSController {
  static bool hasInit = false;
  static late Uri wsUrl;
  static late WebSocketChannel channel;

  static void _judgeNull() {
    if (!hasInit) init();
    print("重新初始化");
  }

  static void init() {
    print("正在初始化 ws   , location : web_socket.dart ,init");
    if (hasInit) return;
    try {
      wsUrl = Uri.parse("${Constant.SOCKET_URL}?userId=${UserController.me.id}");
      channel = IOWebSocketChannel.connect(
        wsUrl,
        pingInterval: const Duration(seconds: 5),
      );

      channel.stream.listen(
        (event) {
          print("收到消息$event");
          try {
            WsReMessage message = WsReMessage.fromJson(json.decode(event));
            switch (message.type) {
              case MessageType.chatMessage:
                ChatResContent content = ChatResContent.fromJson(message.content);
                Get.find<ChatController>().receiveStreamingMessage(content.chatId, content.message);
                break;
              case MessageType.chatSystem:
                break;
              default:
            }
          } catch (e) {
            print("解析消息失败");
          }
        },
        onDone: () {
          hasInit = false;
          print("onDone");
        },
        onError: (e) {
          hasInit = false;
          print("ws 出现错误，尝试重连");
          init();
        },
      );
      hasInit = true;
      print("初始化成功");
    } catch (e) {
      print(e);
      hasInit = false;
    }
  }

  //向服务器发送消息
  static void send(String message) {
    _judgeNull();
    channel.sink.add(message);
  }
}
