#!/usr/bin/env python3
#-*- coding:utf-8 -*-

import tensorflow as tf
from tensorflow.examples.tutorials.mnist import input_data

num = 1000  # 设定训练次数 
print("指定的训练次数是：",num)
# 获取数据
mnist = input_data.read_data_sets("MNIST_data/", one_hot=True)

# 构建图
x = tf.placeholder("float", [None, 784])
W = tf.Variable(tf.zeros([784,10]))
b = tf.Variable(tf.zeros([10]))

y = tf.nn.softmax(tf.matmul(x,W) + b)
y_ = tf.placeholder("float", [None,10])
cross_entropy = -tf.reduce_sum(y_*tf.log(y))
train_step = tf.train.GradientDescentOptimizer(0.01).minimize(cross_entropy)

# 进行训练
init = tf.global_variables_initializer()
sess = tf.Session()
sess.run(init)

for i in range(num): 
  batch_xs, batch_ys = mnist.train.next_batch(100)
  sess.run(train_step, feed_dict={x: batch_xs, y_: batch_ys})

# 模型评估
correct_prediction = tf.equal(tf.argmax(y,1), tf.argmax(y_,1))
accuracy = tf.reduce_mean(tf.cast(correct_prediction, "float"))
print("准确率：",sess.run(accuracy, feed_dict={x: mnist.test.images, y_: mnist.test.labels}))
