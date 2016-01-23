// Sample code for communicating between
// PC <-> USB <-> Arduino <-> ESP8266
// PC is connected to Serial using USB
// ESP8266 is connected to 8 and 9 pins

#include <SoftwareSerial.h>
SoftwareSerial softSerial(8, 9); // RX, TX

void setup() 
{
  uint32_t baud = 9600;
  Serial.begin(baud);
  softSerial.begin(baud);
  Serial.print("SETUP!! @");
  Serial.println(baud);
}

void loop() 
{
    while(softSerial.available() > 0) 
    {
      char a = softSerial.read();
      // Filter away junk characters sometimes read from
      // the software serial
      if(a == '\0')
        continue;
      if(a != '\r' && a != '\n' && (a < 32))
        continue;
      Serial.print(a);
    }
    
    while(Serial.available() > 0)
    {
      char a = Serial.read();
      Serial.write(a);
      softSerial.write(a);
    }
}

