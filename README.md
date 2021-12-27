# crawler

<!-- https://codertw.com/%E7%A8%8B%E5%BC%8F%E8%AA%9E%E8%A8%80/30332/ -->


reference：
https://ithelp.ithome.com.tw/questions/10198403
https://blog.jiatool.com/posts/job104_spider/
https://www.royenotes.com/python-104-employment-agency/

curl --location --request GET 'https://www.104.com.tw/job/6hr41?jobsource=cs_2018indexpoc' \
--header 'Accept-Language: en-US,en;q=0.9,zh-TW;q=0.8,zh;q=0.7' \
--header 'Referer: https://www.104.com.tw/job/6hr41' \
--header 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36' \

<!-- Notes -->
技能會被放在： 工作內容、擅長工具、其他條件 裡面，要再修改