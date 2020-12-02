# coding=utf-8
from matplotlib  import pyplot as plt


x1 = range(2,26,2)
y1 = [12,13,14,15,17,18,19,23,24,23,44,56,78,23]

#create map func
#fig = plt.figure(figsize=(20,8),dpi=80)
plt.figure(figsize=(200,80),dpi=80)
#plt.style("ggplot")


#draw the pric
plt.plot(x1,y1)

#save pricetrue
plt.savefig("./sig_size.png")

#show procetrue
plt.show()



#if __name__ =='__main__':
#    plt.savefig("./sig_size.png")
#    plt.show()

