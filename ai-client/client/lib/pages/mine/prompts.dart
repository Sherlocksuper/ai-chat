import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';

class PromptList extends StatelessWidget {
  PromptList({super.key});

  final List<Map<String, String>> prompts = [
    {
      'title': '标题1',
      'content': '内容1',
    },
    {
      'title': '标题2',
      'content': '内容2',
    },
    // 添加更多prompt...
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Prompts'),
      ),
      body: ListView.builder(
        itemCount: prompts.length,
        itemBuilder: (context, index) {
          return PromptItem(
            title: prompts[index]['title']!,
            content: prompts[index]['content']!,
          );
        },
      ),
    );
  }
}

class PromptItem extends StatelessWidget {
  const PromptItem({
    super.key,
    required this.title,
    required this.content,
  });

  final String title;
  final String content;

  @override
  Widget build(BuildContext context) {
    return Card(
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(10),
      ),
      elevation: 4,
      margin: const EdgeInsets.all(10),
      child: Padding(
        padding: const EdgeInsets.all(10),
        child: Row(
          children: [
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    title,
                    style: const TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 16,
                    ),
                  ),
                  const SizedBox(height: 5),
                  Text(
                    content * 99,
                    style: const TextStyle(fontSize: 14),
                    overflow: TextOverflow.ellipsis,
                    maxLines: 4, // 可以根据需要调整
                  ),
                ],
              ),
            ),
            IconButton(
              icon: const Icon(Icons.copy),
              onPressed: () {
                Clipboard.setData(ClipboardData(text: '$title\n$content'));
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(
                    content: Text('Prompt copied to clipboard!'),
                  ),
                );
              },
            ),
          ],
        ),
      ),
    );
  }
}

class Prompt {
  String id;
  String title;
  String description;
  String content;

  Prompt({
    required this.id,
    required this.title,
    required this.description,
    required this.content,
  });

  //fromJson
  factory Prompt.fromJson(Map<String, dynamic> json) {
    return Prompt(
      id: json['id'],
      title: json['title'],
      description: json['description'],
      content: json['content'],
    );
  }
}
