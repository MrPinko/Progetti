import sys


import pyautogui
from pynput import keyboard

velocity = 2         # 1= fermo
exitProgram = False
pauseProgram = False


# funzione
def on_press(key):
    global velocity, exitProgram, pauseProgram
    if key == keyboard.Key.esc:
        exitProgram = True
    elif key == keyboard.Key.space:
        pauseProgram = not pauseProgram
    elif key.char == ('+'):
        velocity += 1

    elif key.char == ('-') and velocity > 1:
        velocity -= 1

    print(velocity)


listener = keyboard.Listener(on_press=on_press)
listener.start()

print("START")

while True:
    if exitProgram:
        sys.exit()
    if not pauseProgram:
        pyautogui.scroll(-velocity)
