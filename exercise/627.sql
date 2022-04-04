update
    salary
set
    sex = case sex when 'f' then 'm' else 'f' end

update
    salary
set
    sex = replace('fm',sex,'')
/*repalce 新用法，效率更高*/