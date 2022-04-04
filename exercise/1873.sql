select 
    employee_id,
    case when mod(employee_id,2)=1 and name not like 'M%' then salary else 0 end as bonus
from
    employees
order by
    employee_id

/* not like 'M%' 575ms 
    not rlike '^M' 704ms */