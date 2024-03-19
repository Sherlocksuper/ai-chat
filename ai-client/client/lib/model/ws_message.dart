//从ws接收的信息格式
import 'package:get/get.dart';
import '../Controller/chat_controller.dart';

class WsMessage {
  String type = "chat";
  Map content;
  Function? callback;

  WsMessage({
    required this.type,
    required this.content,
  });
}

enum WsFuncType {
  chat,
  system,
}

//收到返回的ChatMessage
void onReceiveChatMessage(Map data) {
  print("onReceiveChatMessage + $data");
  Get.find<ChatController>().addMessage(data["chatId"], "assistant", data["message"]);
}
