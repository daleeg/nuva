class Teacher(object):
    def __init__(self, name):
        self.name = name

    @property
    def data(self):
        return {"name": self.name}

    def display_json(self, visitor):
        print(visitor.to_json(self.data))


class Visitor(object):
    def __init__(self, name):
        self.name = name

    def to_json(self, data):
        import json
        data["visitor"] = self.name
        return json.dumps(data)


def test_visitor():
    t = Teacher("lao wang")
    v = Visitor("json")
    t.display_json(v)


if __name__ == "__main__":
    test_visitor()
