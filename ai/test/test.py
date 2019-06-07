from skimage import transform, io
import argparse


parser=argparse.ArgumentParser()
parser.add_argument("img1", help="The First Image")
parser.add_argument("img2", help="The Second Image")
args=parser.parse_args()

img1=io.imread(args.img1)
img2=io.imread(args.img2)
if sum(img1.shape) > sum(img2.shape):
	scaler=[img1.shape[ind]/img2.shape[ind] for ind in range(2)]
	img2_new=transform.rescale(img2, scaler)
	io.imsave(args.img2, img2_new)
	del img1
	del img2
	del img2_new
else:
	scaler=[img2.shape[ind]/img1.shape[ind] for ind in range(2)]
	img1_new=transform.rescale(img1, scaler)			
	io.imsave(args.img1, img1_new)
	del img1
	del img2
	del img1_new
