select 
    a.id
from
    Weather a,
    Weather b
where
    a.temperature > b.temperature
    and
    datediff(a.recordDate,b.recordDate)=1

/* datediff 计算两个日期的差值！ */