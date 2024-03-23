//从ws接收的信息格式
import 'package:get/get.dart';
import '../Controller/chat_controller.dart';

// Define the message types as a class with static constants.
class MessageType {
  static const String chatMessage = 'chat_message';
  static const String chatSystem = 'chat_system';
}

// Define the WebSocket response message class.
class WsReMessage {
  String type;
  dynamic content; // 'any' in Dart is represented as 'dynamic'.

  WsReMessage({required this.type, this.content});

  // Method to convert a JSON map to a WsReMessage instance.
  factory WsReMessage.fromJson(Map<String, dynamic> json) {
    return WsReMessage(
      type: json['type'],
      content: json['content'],
    );
  }

  // Method to convert a WsReMessage instance to a JSON map.
  Map<String, dynamic> toJson() {
    return {
      'type': type,
      'content': content,
    };
  }
}

// Define the chat response content class.
class ChatResContent {
  int userId;
  int chatId;
  String message;

  ChatResContent({required this.userId, required this.chatId, required this.message});

  // Method to convert a JSON map to a ChatResContent instance.
  factory ChatResContent.fromJson(Map<String, dynamic> json) {
    return ChatResContent(
      userId: json['userId'],
      chatId: json['chatId'],
      message: json['message'],
    );
  }

  // Method to convert a ChatResContent instance to a JSON map.
  Map<String, dynamic> toJson() {
    return {
      'userId': userId,
      'chatId': chatId,
      'message': message,
    };
  }
}
