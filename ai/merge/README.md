Face Merge
===================


This is a tool which applies merge between emoji face and human face.

----------

Requirements
-------------
Install necessary packages and tools by the command below: 
```
pip install -r requirements.txt
```
:bulb: Hints for the installation of **dlib** and **opencv**: 
- For macOS, link [here](https://www.learnopencv.com/install-dlib-on-macos/)   
- For Linux, link [here](https://www.learnopencv.com/install-dlib-on-ubuntu/   
- For Windows, link [here](https://www.learnopencv.com install-opencv-3-and-dlib-on-windows-python-only/)

Usage
-------------

Locate the **merge** folder, run python command as follows:
```
python merge_face.py path_to_first_image(emoji_face) path_to_second_image(human_face) path_to_store_the_merged_face

e.g python merge_face.py ../data/emoji.jpg ../data/source.jpg ../data/merged_face.jpg
```
- Note that no need on the type of input photo. no need on the same shape between two input photos

To-dos
-------------
...

How It Works
-------------
1. Find point-to-point correspondences between the two images using Dlib's **Facial Landmark Detection**.
2. ...

Examples
-------------
![A photo of emoji face]()
![A phtot of human face]()
![Merged face](.jpg)

