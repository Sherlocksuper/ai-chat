class Message {
  int? id;
  String? createdAt;
  String? updatedAt;
  int chatId;
  String role;
  String content;

  Message({
    this.id,
    this.createdAt,
    this.updatedAt,
    required this.chatId,
    required this.role,
    required this.content,
  });

  factory Message.fromJson(Map<String, dynamic> json) {
    return Message(
      id: json['id'],
      createdAt: json['createdAt'],
      updatedAt: json['updatedAt'],
      chatId: json['chatId'],
      role: json['role'],
      content: json['content'],
    );
  }
}
