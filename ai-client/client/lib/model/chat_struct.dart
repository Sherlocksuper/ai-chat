import 'message.dart';

class ChatDetailStruct {
  int id;
  String createdAt;
  String updatedAt;
  String title;
  int userId;
  String systemMessage;
  List<Message> messages;

  ChatDetailStruct({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.title,
    required this.userId,
    required this.systemMessage,
    required this.messages,
  });

  factory ChatDetailStruct.fromJson(Map<String, dynamic> json) {
    return ChatDetailStruct(
      id: json['id'],
      createdAt: json['createdAt'],
      updatedAt: json['updatedAt'],
      title: json['title'],
      userId: json['userId'],
      systemMessage: json['systemMessage'],
      messages: List<Message>.from(json['messages'].map((e) => Message.fromJson(e))),
    );
  }
}
