import requests
import xml.etree.ElementTree as ET

url = "http://webservices.dotnethell.it/codicefiscale.asmx/CalcolaCodiceFiscale"  # host
myobj = {'Nome': '',           
         'Cognome': '',             
         'ComuneNascita': '',       
         'DataNascita': '',  	# formato europeo (richiesto dal sito)  gg-MM-aaaa
         'Sesso': ''}

x = requests.post(url, data=myobj)

with open("respose.xml", "w") as f:
    f.write(x.text)

tree = ET.parse('respose.xml')
root = tree.getroot()

print(root.text)
