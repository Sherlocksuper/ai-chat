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
      id: json['id'] ?? json['ID'],
      createdAt: json['createdAt'] ?? json['CreatedAt'],
      updatedAt: json['updatedAt'] ?? json['UpdatedAt'],
      title: json['title'] ?? json['Title'],
      userId: json['userId'] ?? json['UserId'],
      systemMessage: json['systemMessage'],
      messages: List<Message>.from(json['messages'].map((e) => Message.fromJson(e))),
    );
  }
}
