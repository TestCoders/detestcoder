package com.training;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.util.List;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;

class InMemoryCustomerRepositoryTest {

    private InMemoryCustomerRepository inMemoryCustomerRepository;

    @BeforeEach
    public void setup(){
        inMemoryCustomerRepository = new InMemoryCustomerRepository();
    }

    private void preloadThreeCustomers(){
        Customer customer1 = new Customer(1,"Een", "A");
        Customer customer2 = new Customer(2,"Twee", "B");
        Customer customer3 = new Customer(3,"Drie", "C");
        inMemoryCustomerRepository.save(customer1);
        inMemoryCustomerRepository.save(customer2);
        inMemoryCustomerRepository.save(customer3);
    }

    @Test
    public void shouldGetSpecificCustomerWhenPresent(){
        Customer singleCustomer = new Customer(10,"single","Customer");
        inMemoryCustomerRepository.save(singleCustomer);
        Optional<Customer> customer = inMemoryCustomerRepository.getCustomer(singleCustomer.getId());
        assertTrue(customer.isPresent());
        assertEquals(singleCustomer.getId(),customer.get().getId());
        assertEquals(singleCustomer.getFirstName(),customer.get().getFirstName());
        assertEquals(singleCustomer.getLastName(),customer.get().getLastName());
    }

    @Test
    public void shouldGetOptionalCustomerWhenNoCustomerPresent(){
        Optional<Customer> customer = inMemoryCustomerRepository.getCustomer(66);
        assertFalse(customer.isPresent());
    }

    @Test
    public void shouldGetListOfCustomersWhenMultiplePresent(){
        preloadThreeCustomers();
        List<Customer> customers = inMemoryCustomerRepository.getCustomers();
        assertTrue(customers.size()>1);
    }

    @Test
    public void shouldGetListOfCustomersWhenSinglePresent(){
        Customer singleCustomer = new Customer(15,"single","Customer");
        inMemoryCustomerRepository.save(singleCustomer);
        List<Customer> customers = inMemoryCustomerRepository.getCustomers();
        assertEquals(1, customers.size());
    }

}