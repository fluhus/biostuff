import numpy as np
import matplotlib.pyplot as plt

a = np.loadtxt("a.txt")
y = a[:,0]
x = a[:,2:]

positives = x[y == 1,:]
negatives = x[y == -1,:]

howmany = 100
plt.plot(negatives[:howmany,:].transpose(), "ro")
plt.plot(positives[:howmany,:].transpose(), "b+")
plt.axis([-2, 10, -2, 10])
plt.show()

