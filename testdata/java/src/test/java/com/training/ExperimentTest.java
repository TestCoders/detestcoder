package com.training;

import org.junit.jupiter.api.Test;

import java.util.Optional;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class ExperimentTest {
    @Test
    public void givenOptional_whenFlatMapWorks_thenCorrect2() {
        Person person = new Person("john", 26);
        Optional<Person> personOptional = Optional.of(person);
        System.out.println("personOptional: "+personOptional);

        Optional<Optional<String>> nameOptionalWrapper
                = personOptional.map(Person::getName);
        System.out.println("nameOptionalWrapper: "+nameOptionalWrapper);

        Optional<String> nameOptional
                = nameOptionalWrapper.orElseThrow(IllegalArgumentException::new);
        System.out.println("nameOptional: "+nameOptional);

        String name1 = nameOptional.orElse("");
        System.out.println("name1: "+name1);

        assertEquals("john", name1);

        // Flatmap
        Optional<String> nameFM = personOptional
                .flatMap(Person::getName);
        System.out.println("FM String: "+nameFM);


        String name = personOptional
                .flatMap(Person::getName)
                .orElse("");
        assertEquals("john", name);
    }

}
