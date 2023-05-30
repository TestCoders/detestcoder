package com.training;

import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.Mockito;

import java.util.Optional;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.ArgumentMatchers.any;
import static org.mockito.ArgumentMatchers.anyInt;
import static org.mockito.Mockito.*;

class CustomerServiceTest {
//    @Mock
//    private CustomerRepository customerRepository;
//
//    @InjectMocks
//    private CustomerService customerService;

    private CustomerRepository customerRepository;
    private CustomerService customerService;

    @BeforeEach
    public void setup(){
        customerRepository = Mockito.mock(CustomerRepository.class);
//        customerRepository = new InMemoryCustomerRepository();
        customerService = new CustomerService(customerRepository);
    }

    @Test
    public void shouldReturnOptionalEmptyWhenCustomerNotFound(){
        when(this.customerRepository.getCustomer(anyInt())).thenReturn(Optional.empty());
        Optional<Customer> customer = customerService.getCustomer(2);
        assertEquals(Optional.empty(), customer);
    }

    @Test
    public void shouldReturnCustomerWhenFound(){
        Customer testCustomer = new Customer(5,"Gerwin", "Vaatstra");
        when(this.customerRepository.getCustomer(5)).thenReturn(Optional.of(testCustomer));
        Optional<Customer> customer = customerService.getCustomer(5);
        assertNotEquals(Optional.empty(), customer);
        assertEquals("Gerwin", customer.get().getFirstName());
        assertEquals("Vaatstra", customer.get().getLastName());
        assertEquals(5, customer.get().getId());
    }

    @Test
    public void shouldSaveCustomer(){
        Customer myTestCustomer = new Customer(3,"A","B");
        customerService.save((myTestCustomer));
        verify(customerRepository, times(1)).save(any(Customer.class));
    }

}