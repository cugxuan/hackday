Face Morphing
===================


This is a tool which creates a morphing effect. It takes two facial images as input and returns a video morphing from the first image to the second.

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

Besides packages, software **FFMPEG** is necessary for the video production, download link is [here](https://ffmpeg.zeranoe.com/builds/). No more step shall be taken except unzipping. Remember to set enviroment variable for the executable program.

Usage
-------------

Locate the **morphing** folder, run python command as follows:
```
python run_face_morph.py face/body path_to_first_image path_to_second_image duration_of_video frame_rate path_where_video_is_to_be_saved

e.g. python run_face_morph.py face ../data/source.jpg ../data/merged_face.jpg 5 25 ../data/morphing-example.mp4
```
- Note that presently only **face** is allowed as first argument.
- Meanwhile, only two photos, which share almost same shape, sucess up to now. Otherwise, resize by opencv package may fail in the fubnction "doCropping" from [face_landmark_detection.py](./face_landmark_detection.py).
- Instead of the path to image file, Python type: **File Object** of that image can also be passed to the command. 

To-dos
-------------
2019/06/07: fix the requirement for the shape og two input photos
2019/06/07: photo file path use "file:///D:/Download/Deemo2/demos/ben.jpg" rather than "D:/Download/Deemo2/demos/ben.jpg", why?

How It Works
-------------
1. Find point-to-point correspondences between the two images using Dlib's **Facial Landmark Detection**.
2. Find the **Delaunay Triangulation** for the average of these points.
3. Using these corresponding triangles in both initial and final images, perform **Warping** and **Alpha Blending** and obtain intermediate images to be used in creating videos.
4. Use **ffmpeg** to return a video from above created frames.

Examples
-------------
![A photo of raw human face]()
![A phtot of merged face]()
![Morphed Video](.gif)

