class Message {
  int? id = 0;
  String? createdAt = DateTime.now().toString();
  String? updatedAt = DateTime.now().toString();
  int? chatId = 0;
  String role;
  String content;

  Message({
    this.id,
    this.createdAt,
    this.updatedAt,
    this.chatId,
    required this.role,
    required this.content,
  });

  factory Message.fromJson(Map<String, dynamic> json) {
    return Message(
      id: json['id'] ?? 0,
      createdAt: json['createdAt'] ?? DateTime.now().toString(),
      updatedAt: json['updatedAt'] ?? DateTime.now().toString(),
      chatId: json['chatId'],
      role: json['role'],
      content: json['content'],
    );
  }
}
