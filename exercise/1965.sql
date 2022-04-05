-- Write your MySQL query statement below
select 
    a.employee_id
from
    employees a
left join
    salaries b
on
    a.employee_id = b.employee_id
where 
    b.salary is null

union

select 
    a.employee_id
from
    salaries a 
left join  
    employees b
on
    a.employee_id = b.employee_id
where
    b.name is null

order by employee_id

/* 这个快一点 */
select employee_id from employees
where employee_id not in (select employee_id from salaries)

union

select employee_id from salaries
where employee_id not in(select employee_id from employees)

order by employee_id

/*union all */
select employee_id from(
    select employee_id from employees
    union all
    select employee_id from salaries
)as tmp
group by employee_id
having count(employee_id) = 1
order by employee_id