import 'dart:async';

import 'package:client/model/chat_struct.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_markdown/flutter_markdown.dart';
import 'package:get/get.dart';
import '../../Controller/ChatController.dart';
import '../../model/message.dart';

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
        return FutureBuilder(
          future: () async {}(),
          builder: (context, snapshot) {
            return Scaffold(
              appBar: AppBar(
                title: Text(chat.title),
              ),
              body: Column(
                children: [
                  Expanded(
                    child: ListView.builder(
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
      //如果role为user，在右面、role为system在中间，role为other在左面
      mainAxisAlignment: role == "user"
          ? MainAxisAlignment.end
          : role == "system"
              ? MainAxisAlignment.center
              : MainAxisAlignment.start,
      children: [
        Container(
          constraints: BoxConstraints(maxWidth: MediaQuery.of(context).size.width * 0.6),
          decoration: BoxDecoration(
            color: role == "user" ? Colors.green : Colors.blue,
            borderRadius: BorderRadius.circular(10),
          ),
          padding: const EdgeInsets.all(10),
          margin: const EdgeInsets.symmetric(horizontal: 10, vertical: 5),
          child: MarkdownBody(
            data: message,
            selectable: true,
            styleSheet: MarkdownStyleSheet(
              p: const TextStyle(color: Colors.white),
              code: const TextStyle(color: Colors.white, backgroundColor: Colors.black),
              codeblockDecoration: const BoxDecoration(
                color: Colors.black,
              ),
            ),
          ),
        ),
      ],
    );
  }
}
