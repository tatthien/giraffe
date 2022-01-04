---
title: Cấu hình Golang web application với Nginx
date: 2016-11-09T09:08:00.698Z
tags: golang,nginx,web
---

Mấy bữa nay mình có code một web app bằng Golang để cung cấp API (phục vụ cho mục đích cá nhân). Mọi thứ ở local đã xong, việc mình cần làm là đưa lên host để có thể sử dụng live được.

Cái khó là việc cấu hình này lần đầu mình làm nên cũng hơi gian nan. Tuy nhiên sau vài ngày thì mình cũng làm xong. Cho nên viết ngay bài này để note lại lỡ sau này có quên thì còn vô đây coi.

Phần này mình chia làm 3 phần như sau:

1.  **Chạy web app dạng deamon** (run in background)
2.  **Thiết lập cron job** để tự động lấy dữ liệu theo thời gian quy định (vì API này là crawl tin tức mà :smile:)
3.  **Cấu hình proxy** cho web app trên nginx.

## Chạy web app dạng deamon.

Mỗi lần muốn run server thì mình phải `go run main.go` lúc đó nó sẽ chạy 1 server ở địa chỉ là `localhost:2508`. 

Tuy nhiên nếu bấm `Ctrl + C` thì sẽ bị ngắt liền, nên mình cần phải chạy web app này dạng deamon (tức là chạy ngầm).

Cách làm này mình làm như sau:

```bash
go run build main.go
nohup ./main &
```

Để kiểm tra nó có chạy chưa thì bạn có thể dùng lệnh sau:

```bash
ps aux | grep main
```

Sau khi chạy lệnh đó bạn có thể thấy được `PID` và muốn tắt thì có thể dùng lệnh sau:

```bash
kill -9 <PID>
```

## Thiết lập cron job

Cron job hiểu đơn giản là hệ thống sẽ tự động thực thi một cái gì đó vào thời điểm nhất định ví dụ như tự động gởi mail vào thứ 2 hàng tuần.

Đối với app của mình thì mình sẽ thiết lập tự động crawl dữ liệu sau mỗi 20 phút.

Để thêm cron job bạn gõ lệnh:

```bash
crontab -e
```

Sau đó thêm dòng lệnh thực thi như sau:

```bash
*/20 * * * * cd /path/to/source && ./file_to_execute > /dev/null 2>&1
```

Các bạn có thể [vào đây](https://crontab.guru/) để convert syntax trên thành human readable.

Để xem lại danh sách các cronjob thì bạn gõ lệnh:

```bash
crontab -l
```

Đọc thêm về cron job [tại đây](http://www.cyberciti.biz/faq/how-do-i-add-jobs-to-cron-under-linux-or-unix-oses/).

## Cấu hình proxy

Phần này mình không chắc gọi là cấu hình proxy có đúng không, vì mình chưa rành nginx lắm cũng như mình không nghiên cứu khái niệm của nó là gì, chỉ tìm cách giải quyết được vấn để của mình là làm sao có thể truy cập vào domain `api.mydomain.com/app/` thì nó hiểu được là mình đang truy cập vào `localhost:2508.`

Cách làm của mình như sau:

Đầu tiên mình mở file config của domain api.mydomain.com là `api.mydomain.com.conf`

```bash
sudo vi api.mydomain.com.conf
```

tiếp theo mình thêm đoạn code config này vào:

```bash
location /app/ {
  proxy_set_header X-Real-IP $remote_addr;
  proxy_set_header X-Forwarded-For $remote_addr;
  proxy_set_header Host $host;
  proxy_pass http://127.0.0.1:2508/;
}
```

Cuối cùng là lưu lại và restart nginx:

```bash
systemctl restart nginx.service
```

Như vậy là xong, bây giờ thay vì truy cập vào `my_ip:2508/api/v1` thì mình sẽ vào `api.mydomain.com/app/api/v1`.

## Kết luận

Bài viết này mình viết lại chủ yếu là để sau này có thể coi lại nếu quên. Trong quá trình đọc nếu bạn thấy có chỗ nào không đúng thì có thể comment cho mình biết với nhé. Thank you!
