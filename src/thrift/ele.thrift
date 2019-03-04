namespace go ele

struct Student{
  1: i32 id,
  2: string name,
  3: string duty,
  4: string department
}

const map<string, string> MAPCONSTANT = {'name': '饿了么大前端', 'group': 'Alibaba'}

service eleThrift {
  Student findStuById(1: i32 id),
  Student findStuByName(1: string name),
  list<Student> findStuByDuty(1: string duty)
  void put(1: Student stu)
}