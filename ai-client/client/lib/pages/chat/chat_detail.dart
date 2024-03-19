import 'package:client/model/chat_struct.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_markdown/flutter_markdown.dart';
import 'package:get/get.dart';
import '../../Controller/chat_controller.dart';

class ChatDetail extends StatelessWidget {
  final ChatDetailStruct chat;
  final textInputController = TextEditingController();

  ChatDetail({super.key, required this.chat});

  final ScrollController _scrollController = ScrollController();

  @override
  Widget build(BuildContext context) {
    return GetBuilder<ChatController>(
      id: chat.id,
      builder: (logic) {
        return Scaffold(
          appBar: AppBar(
            title: Text(chat.title),
          ),
          body: Column(
            children: [
              Expanded(
                child: Stack(
                  children: [
                    ListView.builder(
                      shrinkWrap: true,
                      physics: const ClampingScrollPhysics(),
                      controller: _scrollController,
                      itemBuilder: (context, index) {
                        return ChatBubbles(
                          message: chat.messages[index].content,
                          role: chat.messages[index].role,
                        );
                      },
                      itemCount: chat.messages.length,
                    ),
                    Positioned(
                      top: 10,
                      right: 10,
                      child: FloatingActionButton(
                        onPressed: () {
                          _scrollController.animateTo(
                            _scrollController.position.maxScrollExtent,
                            duration: const Duration(milliseconds: 300),
                            curve: Curves.easeOut,
                          );
                        },
                        mini: true,
                        child: const Icon(Icons.arrow_downward),
                      ),
                    ),
                  ],
                ),
              ),
              Container(
                height: 50,
                padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
                decoration: BoxDecoration(
                  color: Colors.grey[200], // Light grey color for the background
                  boxShadow: [
                    BoxShadow(
                      color: Colors.grey.withOpacity(0.5), // Shadow color with some transparency
                      spreadRadius: 2,
                      blurRadius: 4,
                      offset: const Offset(0, 3), // Position of the shadow
                    ),
                  ],
                ),
                child: Row(
                  children: <Widget>[
                    Expanded(
                      child: TextField(
                        controller: textInputController,
                        decoration: InputDecoration(
                          contentPadding: const EdgeInsets.symmetric(vertical: 5.0, horizontal: 10.0),
                          hintText: '输入消息...',
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(20.0), // Rounded corners for the input field
                          ),
                          fillColor: Colors.white,
                          filled: true,
                        ),
                        onChanged: (value) {
                          textInputController.text = value;
                        },
                      ),
                    ),
                    IconButton(
                      icon: const Icon(Icons.send, color: Colors.blue), // Icon color changed to blue
                      onPressed: () async {
                        var content = textInputController.text;
                        textInputController.clear();

                        await logic.sendMessage(chat.id, content);

                        _scrollController.animateTo(
                          _scrollController.position.maxScrollExtent,
                          duration: const Duration(milliseconds: 300),
                          curve: Curves.easeOut,
                        );
                      },
                    ),
                  ],
                ),
              )
            ],
          ),
        );
      },
    );
  }
}

class ChatBubbles extends StatelessWidget {
  final String message;
  final String role;

  const ChatBubbles({super.key, required this.message, required this.role});

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: role == "user"
          ? MainAxisAlignment.end
          : role == "system"
              ? MainAxisAlignment.center
              : MainAxisAlignment.start,
      children: [
        Container(
          constraints: BoxConstraints(maxWidth: MediaQuery.of(context).size.width * 0.6),
          decoration: BoxDecoration(
            color: _getBubbleColor(role),
            borderRadius: BorderRadius.circular(10),
            boxShadow: [
              BoxShadow(
                color: Colors.black.withOpacity(0.1),
                spreadRadius: 1,
                blurRadius: 3,
                offset: Offset(0, 1),
              ),
            ],
          ),
          padding: const EdgeInsets.all(10),
          margin: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
          child: MarkdownBody(
            data: message,
            selectable: true,
            styleSheet: MarkdownStyleSheet(
              p: TextStyle(color: Colors.white),
              code: TextStyle(color: Colors.white, backgroundColor: Colors.grey[850]),
              codeblockDecoration: BoxDecoration(
                color: Colors.grey[850],
                borderRadius: BorderRadius.circular(8),
              ),
            ),
          ),
        ),
      ],
    );
  }

  Color _getBubbleColor(String role) {
    switch (role) {
      case 'user':
        return Color(0xFF1DA8B8); // Placid Sea
      case 'system':
        return Color(0xFFCCC6BE); // Light Gray
      default:
        return Color(0xFFE39C32); // Intense Yellow
    }
  }
}
