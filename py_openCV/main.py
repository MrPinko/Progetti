import cv2
import numpy as np
import argparse
import imutils
from matplotlib import pyplot as plt

# img = cv2.imread('py_openCV/watch.jpg')
img = cv2.imread('py_openCV/oggetti.jpg')

(h, w, d) = img.shape
rotation_grade = 0

# * ridimensionamento dell'immagine senza distorcerla
# r = 500.0 / w
# dim = (500, int(h * r))
# resized = cv2.resize(img, dim)
# cv2.imshow("Aspect Ratio Resize", resized)


# * rotazione
# center = (w // 2, h // 2)
# M = cv2.getRotationMatrix2D(center, rotation_grade, 1.0)
# rotated = cv2.warpAffine(img, M, (w, h))
# cv2.imshow('image', rotated)

# *rotazione con uno slider
# def nothing(x):
#     global rotation_grade
#     rotation_grade = x
#     drawImage()

# def drawImage():
#     global rotation_grade
#     global img
#     (h, w, d) = img.shape
#     center = (w // 2, h // 2)
#     M = cv2.getRotationMatrix2D(center, rotation_grade, 1.0)
#     rotated = cv2.warpAffine(img, M, (w, h))
#     cv2.imshow('image', rotated)

# # controlli slider
# cv2.namedWindow('controls')
# cv2.createTrackbar('r', 'controls', 0, 360, nothing)
# cv2.imshow('image', img)


# * riconoscimento oggetti
# edged = cv2.Canny(img, 30, 150)
# cv2.imshow("Edged", edged)
# cv2.imshow("normal", img)

# * applicare thrash immagine in bianco e nero per distinguere lo sfondo dagli oggetti
# gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
# thresh = cv2.threshold(gray, 225, 255, cv2.THRESH_BINARY_INV)[1]
# cv2.imshow("Thresh", thresh)

# * riconoscimento dei contorni
# gray = cv2.cvtColor(img, cv2.COLOR_BGR2GRAY)
# thresh = cv2.threshold(gray, 225, 255, cv2.THRESH_BINARY_INV)[1]
# cnts = cv2.findContours(thresh.copy(), cv2.RETR_EXTERNAL,
#                         cv2.CHAIN_APPROX_SIMPLE)
# cnts = imutils.grab_contours(cnts)
# output = img.copy()

# text = "trovati {} oggetti".format(len(cnts))
# cv2.putText(output, text, (10, 25),
#             cv2.FONT_HERSHEY_SIMPLEX, 0.7, (249, 0, 159), 2)

# for c in cnts:
#     cv2.drawContours(output, [c], -1, (240, 0, 159), 3)
#     cv2.imshow("Contours", output)
#     cv2.waitKey(0)


cv2.waitKey(0)
cv2.destroyAllWindows()
