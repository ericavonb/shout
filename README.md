# SHOUT!
Print large block letters. Because when you are frustrated with Go, you need your terminal output to really scream.

## How to use
To use,
`` import "https://github.com/ericavonb/shout"``

To make a string into block letters, use 'Blockify':
``Blockify("BOO!", 80)``
returns:
" ______       ____       ____     _     
|\ _____\   /\ ____\   /\ ____\  |\_\   
| |  __  | | /  ___  \| /  ___  \| | |  
| | |_\| / | | | | | || | | | | || | |  
| |  __  \ | | |_| | || | |_| | | \|_|  
| | |_\| | \ | |__\| |\ | |__\| ||\_\   
 \|______/   \______/   \______/  \|_|  "
 
 Blockify takes a string and the maximum width.
 
 To print out your message, use 'PrintBlocks':
 
 ``PrintBlocks("No go!") ``
 prints:
 "__    __     ____             ______       ____     _    
|\ _\ |\ \  /\ ____\          |\ _____\   /\ ____\  |\_\   
| |  \| | || /  ___  \        | |  ____| | /  ___  \| | |  
| |   \ | || | | | | |        | | |  __  | | | | | || | |  
| | |\ \| || | |_| | |        | | |_|\_\ | | |_| | | \|_|  
| | | \   |\ | |__\| |        | | |__\| |\ | |__\| ||\_\   
 \|_|  \__|  \______/          \|______/   \______/  \|_|  "

'PrintBlocks' uses the width of your terminal as the maximum width.

The currently supported characters are all letters , ".", "-", "_", "?", "!"

## Built in functions
Some built-in functions:
- 'MegaFail(err)` prints "FAIL!" if err is not nil.
- 'BadGo()` prints "BAD GO!"


## Coming soon:
Numbers, more symbols, underlining, colors!	
	
