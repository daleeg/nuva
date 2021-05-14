class Student(object):
    def __init__(self, name):
        self.name = name

    @property
    def data(self):
        return {"name": self.name}


class Display(object):
    def __init__(self, name):
        self.name = name

    def display_json(self, data):
        import json
        data["view"] = self.name
        print(json.dumps(data))


class Control(object):
    def __init__(self, m, v):
        self.m = m
        self.v = v

    def display_student(self):
        self.v.display_json(self.m.data)


def test_mvc():
    student = Student("xiao ming")
    display = Display("json")
    control = Control(student, display)
    control.display_student()

if __name__ == "__main__":
    test_mvc()
