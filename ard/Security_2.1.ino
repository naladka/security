#include <SPI.h>
#include <RFID.h> // библиотека "RFID".
//#include <Wire.h> 
//#include <LiquidCrystal_I2C.h>


#define SS_PIN 53
#define RST_PIN 5
#define dooropen 22 //концевик двери
#define ring 23     //звонок
#define lock 24     //открытие замка снаружи
#define lockopen 25 //команда на открытие замка

//LiquidCrystal_I2C lcd(0x27,16,2);
RFID rfid(SS_PIN, RST_PIN); 

int serNum0;      //переменные RFID
int serNum1;
int serNum2;
int serNum3;
int serNum4;
int incomingByte = 0;   // переменная для хранения полученного байта

int total;
int s0;
int s1;
int s2;
int s3;
int s4;
int changed;

bool door_status = false;
int door = 0;

void setup() 
  {  
  pinMode(dooropen,INPUT);
  pinMode(ring,INPUT);
  pinMode(lock,INPUT);
  pinMode(lockopen,OUTPUT);
  Serial.begin(115200);
//  lcd.init();
//  lcd.backlight();
  SPI.begin();     //  инициализация SPI / Init SPI bus.
  rfid.init();     // инициализация MFRC522 / Init MFRC522 card.
  }


void loop() 
  {
    delay(100);
     while (Serial.available() == 0) {  //если есть доступные данные/
      if (rfid.isCard()) {
              if (rfid.readCardSerial()) {
                  if (rfid.serNum[0] != serNum0
                      && rfid.serNum[1] != serNum1
                      && rfid.serNum[2] != serNum2
                      && rfid.serNum[3] != serNum3
                      && rfid.serNum[4] != serNum4
                  ) {
                      serNum0 = rfid.serNum[0];
                      serNum1 = rfid.serNum[1];
                      serNum2 = rfid.serNum[2];
                      serNum3 = rfid.serNum[3];
                      serNum4 = rfid.serNum[4];
                      changed = 1;
                  }
              }
          } 
      
           
          int DO = digitalRead(dooropen);
          int R = digitalRead(ring);
          int L = digitalRead(lock);
          
          if (DO == 1) {
           serNum0 = 0;
           serNum1 = 0;
           serNum2 = 0;
           serNum3 = 0; 
           serNum4 = 0;
           door_status = true;
          }
           if (door_status == true & DO == 0) {
             door = 1;    
//             Serial.println("true"); 
             door_status = false;
           }
           
           if (((DO != 0 | R != 0 | L != 0 | door == 1)) | ((serNum0 != s0 | serNum1 != s1 | serNum2 != s2 | serNum3 != s3 | serNum4 != s4))){
              Serial.print(DO+2*R+4*L+door*8);
              Serial.print(serNum0+100);
              Serial.print(serNum1+100);
              Serial.print(serNum2+100);
              Serial.print(serNum3+100);
              Serial.print(serNum4+100);
             if (door == 1){
              door = 0;
             }
              
//              lcd.clear();
//              lcd.setCursor(3,0);
//              lcd.print(DO+2*R+4*L);
//              lcd.print(DO+2*R+4*L);
            }
      
            if (!serNum0 == 0)
            {
//              lcd.setCursor(0,1);
//              lcd.print(serNum0+100);
//              lcd.setCursor(3,1);
//              lcd.print(serNum1+100);
//              lcd.setCursor(6,1);
//              lcd.print(serNum2+100);
//              lcd.setCursor(9,1);
//              lcd.print(serNum3+100);
//              lcd.setCursor(12,1);
//              lcd.print(serNum4+100);
            }      


               s0 = serNum0;
               s1 = serNum1;
               s2 = serNum2;
               s3 = serNum3;
               s4 = serNum4;
     }
        incomingByte = Serial.read();
//        lcd.setCursor(6,0);
//        lcd.print(incomingByte);
        if (char(incomingByte)== char(49)) {
             digitalWrite(lockopen, HIGH);
        } else if (char(incomingByte)== char(48)) {
             digitalWrite(lockopen, LOW);
             
         }

}
