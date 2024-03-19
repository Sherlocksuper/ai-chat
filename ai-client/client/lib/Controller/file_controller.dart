import 'dart:io';
import 'package:path_provider/path_provider.dart';

class FileController {
  static String? localPath;

  static bool hasInit = false;

  static Future<void> init() async {
    final directory = await getApplicationDocumentsDirectory();
    localPath = directory.path;
    hasInit = true;
  }

  //向文件写入数据
  static Future<void> writeToFile(String fileName, String data) async {
    if (!hasInit) await init();
    try {
      final file = File('$localPath/$fileName');
      await file.writeAsString(data);
    } catch (e) {
      print(e);
    }
  }

  //从文件读取数据
  static Future<String> readFromFile(String fileName) async {
    if (!hasInit) await init();
    try {
      final file = File('$localPath/$fileName');
      return await file.readAsString();
    } catch (e) {
      print(e);
      return '';
    }
  }

  //清空文件数据
  static Future<bool> clearFileData(String fileName) async {
    if (!hasInit) await init();
    try {
      final file = File("$localPath/$fileName");
      await file.writeAsString("");
    } catch (e) {
      print("clear file data error : $e");
    }
    return true;
  }
}
