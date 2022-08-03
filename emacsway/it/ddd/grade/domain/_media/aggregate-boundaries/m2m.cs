public class Student : Entity
{
    public string Name { get; }
    public string Email { get; }

    private readonly IList<StudentInstructor> _studentInstructors;
    public IReadOnlyList<Instructor> Instructors => _studentInstructors
        .Select(x => x.Instructor)
        .OrderBy(x => x.DateAdded)
        .ToList();

    internal void AddInstructor(StudentInstructor instructor)
    {
        _studentInstructors.Add(instructor);
    }
}

public class Instructor : Entity
{
    public string Name { get; }

    private readonly IList<StudentInstructor> _studentInstructors;
    public IReadOnlyList<Student> Students => _studentInstructors
        .Select(x => x.Student)
        .OrderBy(x => x.DateAdded)
        .ToList();

    public void AddStudent(Student student)
    {
        var studentInstructor = new StudentInstructor(student, this, DateTime.Now);
        _studentInstructors.Add(studentInstructor);
        student.AddInstructor(studentInstructor);
    }
}

public class StudentInstructor : Entity
{
    public Student Student { get; }
    public Instructor Instructor { get; }
    public DateTime DateAdded { get; }

    public StudentInstructor(Student student, Instructor instructor, DateTime dateAdded)
    {
        Student = student;
        Instructor = instructor;
        DateAdded = dateAdded;
    }
}
