# 🔢 baht – ตัวอย่างการใช้งานไลบรารีแปลงจำนวนเงินเป็นข้อความบาทไทย

โปรเจกต์นี้เป็นแอปตัวอย่างที่เขียนด้วยภาษา Go สำหรับสาธิตการใช้งานไลบรารี [`bahttext`](https://github.com/AnuchitO/bahttext) ซึ่งสามารถแปลงค่าตัวเลขเป็นข้อความในรูปแบบภาษาไทย สำหรับจำนวนเงินบาทและสตางค์ — คล้ายกับฟังก์ชัน `BAHTTEXT()` ที่อยู่ใน Microsoft Excel

 <img src=".badges/trustworthy.png" alt="Trustworthy" height="140"> <img src=".badges/production-proved.png" alt="Trustworthy" height="140">

---

## 📦 ต้นทางไลบรารี

👉 ไลบรารีต้นฉบับพัฒนาโดยคุณ [AnuchitO](https://github.com/AnuchitO)
📍 เยี่ยมชมและให้ดาวได้ที่: [github.com/anuchito/bahttext](https://github.com/anuchito/bahttext)

---

## 🚀 วิธีติดตั้งและใช้งาน

```bash
git clone https://github.com/anuchito/baht
cd baht
go mod tidy
go run demo.go
```

ตัวอย่าง
```go
	fmt.Println(bahttext.Words(1234.56))                  // หนึ่งพันสองร้อยสามสิบสี่บาทห้าสิบหกสตางค์
	fmt.Println(bahttext.Words(-100))                     // ลบหนึ่งร้อยบาทถ้วน
	fmt.Println(bahttext.WordsFromString("1,000,000.01")) // หนึ่งล้านบาทหนึ่งสตางค์

	fmt.Println(bahttext.Words(1_000_000.01))     // หนึ่งล้านบาทหนึ่งสตางค์
	fmt.Println(bahttext.Words(1_000_000_000.01)) // หนึ่งล้านบาทหนึ่งสตางค์
```

ตัวอย่างฉบับเต็ม [click](https://github.com/AnuchitO/bahttext?tab=readme-ov-file#%E0%B8%95%E0%B8%B2%E0%B8%A3%E0%B8%B2%E0%B8%87%E0%B8%95%E0%B8%B1%E0%B8%A7%E0%B8%AD%E0%B8%A2%E0%B9%88%E0%B8%B2%E0%B8%87%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%87%E0%B8%B2%E0%B8%99--examples-table)
