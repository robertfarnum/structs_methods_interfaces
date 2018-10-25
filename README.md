# Overview
This Structs, methods, & interfaces project is derived from 
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces

# Scope Creep
This project goes beyond the original scope as outlined in the site above as it explores an alternative specification of the interface design and package layout of the original project.

In the original project the interface was called 'Shape'.  The association drawn from the interface name 'Shape' is that a Rectangle, Circle, and Triangle are types of Shape(s).  Generally speaking this is attempting to use the object oriented concept of inheritance to support the polymorphic behaviours for calculating the area and perimeter of a Shape. So the natural conclusion of this excersise would be to have a calculateArea and calculatePermiter methods of the 'Shape' interface.

As it turns out, based on the definition of a Trianble, this become problematic when it comes time to implement the calculatePerimeter method for a Triangle.  Given only the Base and Height of a Triangle it is not possible to calculate the perimeter of a Triangle.  But, by definition all the methods of the Shape interface must be implemented whenever an implicit binding to the interface is established via the association of the method declaration for a base type such a a Rectangle, Circle, or Triangle with the interface declared methods.

So in the end it became clear that the deficiency was in the interface declaration which needed to be broken down into two seperate interfaces.  One for Area (AreaCalculator) and the other for Perimeter (PerimeterCalculator).  This allow the implementers to pick and choose the interfaces they want to implement.  In the case of a Rectangle and Circle both the CalculateArea and CalculatePerimeter methods are implemented and therefore adhere to both interface specifications.  But, for a Triangle only the CalculateArea method is implemented and therefore only adheres to the AreaCalulator interface specification.