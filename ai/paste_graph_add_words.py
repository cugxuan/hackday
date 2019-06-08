import cv2
from PIL import ImageFont, ImageDraw, Image
import sys, os, math
import numpy as np
import qrcode


def resize_rate(path, fx, fy):
    image = cv2.imread(path)
    im_resize = cv2.resize(image, None, fx=fx, fy=fy)
    cv2.imwrite(path, im_resize)


qr = qrcode.QRCode(
    version=1,
    error_correction=qrcode.constants.ERROR_CORRECT_L,
    box_size=10,
    border=4,
)
qr.add_data(data="https://www.baidu.com")
qr.make(fit=True)
img = qr.make_image()
img.save(sys.argv[2])

src = Image.open(sys.argv[2])
dst = Image.open(sys.argv[3])
box = (4165, 1730)
src.paste(dst, box)
src.save(sys.argv[1])
del src
del dst

text_size1 = 180
text_size2 = 200
img = cv2.imread(sys.argv[2])
img_pil = Image.fromarray(img)
draw = ImageDraw.Draw(img_pil)

# 粘贴文章摘要文字
color1 = (192, 192, 192)
fontpath1 = "simsun.ttc"
text1 = sys.argv[4].split(" ")
font1 = ImageFont.truetype(fontpath1, text_size1)
# loc_init = (633, 1466)
iter = 5
for ind, text in enumerate(text1):
    loc = (633, 1466+ind*text_size1+(ind+1)*iter)
    draw.text(loc, text, font=font1, fill=color1)
# draw.text(loc_init, text1, font=font1, fill=color1)
# 粘贴用户评论文字
color2 = (0, 0, 0)
fontpath2 = "/Users/xuan/UseOnce/hackday/ai/msyh.ttc"
text2 = sys.argv[5]
font2 = ImageFont.truetype(fontpath2, text_size2)
# loc_init = (2209, 375)
iter1, iter2 = 5, 4 * text_size2
for ind, text in enumerate(text1):
    loc = (2209+(ind+1)*iter2, 375+ind*text_size1+(ind+1)*iter)
    draw.text(loc, text, font=font1, fill=color1)
# draw.text(loc_init, text2, font=font2, fill=color2)
img_upd = np.array(img_pil)
cv2.imwrite(sys.argv[1], img_upd)

size = float(os.path.getsize(sys.argv[2])) / 1024 / 1024
expected_size = 0.5  # unit: mb
while size > expected_size:
    rate = math.ceil((size / expected_size) * 10) / 10 + 0.1
    rate = math.sqrt(rate)

    rate = 1.0 / rate
    resize_rate(sys.argv[2], rate, rate)
    size = float(os.path.getsize(sys.argv[2])) / 1024 / 1024
