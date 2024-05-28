import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'package:image/image.dart' as img;
import 'dart:io';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: HomePage(),
    );
  }
}

class HomePage extends StatefulWidget {
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  File? _image;
  File? _imageWithoutBackground;
  final picker = ImagePicker();

  Future<void> _pickImage() async {
    final pickedFile = await picker.pickImage(source: ImageSource.gallery);

    if (pickedFile != null) {
      setState(() {
        _image = File(pickedFile.path);
      });
    }
  }

  Future<void> _removeBackground() async {
    if (_image == null) return;

    // Load the image using the image package
    final imageBytes = await _image!.readAsBytes();
    img.Image image = img.decodeImage(imageBytes)!;

    // Define the background color to remove (e.g., white)
    const int backgroundColor = 0xFFFFFFFF;

    // Iterate over pixels and set background color pixels to transparent
    for (int y = 0; y < image.height; y++) {
      for (int x = 0; x < image.width; x++) {
        if (image.getPixelSafe(x, y) == backgroundColor) {
          image.setPixel(x, y, img.getColorRgba(0, 0, 0, 0));
        }
      }
    }

    // Save the image with transparent background to a new file
    final outputPath = _image!.path.replaceAll('.jpg', '_nobg.png');
    final outputImage = File(outputPath)
      ..writeAsBytesSync(img.encodePng(image));

    setState(() {
      _imageWithoutBackground = outputImage;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Remove Background App'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            _image != null
                ? Image.file(_image!)
                : Text('Nenhuma imagem selecionada.'),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: _pickImage,
              child: Text('Selecionar Imagem'),
            ),
            SizedBox(height: 20),
            ElevatedButton(
              onPressed: _removeBackground,
              child: Text('Remover Fundo'),
            ),
            SizedBox(height: 20),
            _imageWithoutBackground != null
                ? Image.file(_imageWithoutBackground!)
                : Text('Nenhuma imagem sem fundo.'),
          ],
        ),
      ),
    );
  }
}
