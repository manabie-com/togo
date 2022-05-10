package com.todo;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.gson.Gson;
import com.google.gson.JsonObject;
import net.bytebuddy.utility.RandomString;
import org.hamcrest.Matchers;
import org.junit.jupiter.api.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.test.web.servlet.MvcResult;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;

import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@RunWith(SpringRunner.class)
@SpringBootTest
@AutoConfigureMockMvc
class TodoApplicationTests {
    Gson gson = new Gson();
    String token ;

    @Autowired
    MockMvc mockMvc;
    @Autowired
    ObjectMapper mapper;

    @Test
    void contextLoads() {
    }

    @Test
    public void register_success() throws Exception {
         mockMvc.perform(MockMvcRequestBuilders.post("/register").content("{\"username\" : \"123@abc\"," +
                        "\"password\" : \"pass\"," +
                        "\"taskLimit\" : 4" +
                        "}"))
                .andExpect(status().isOk());
    }

    @Test
    public void authenticate_success() throws Exception {
        MvcResult result = mockMvc.perform(MockMvcRequestBuilders.post("/authenticate")
                        .content("{\"username\" : \"123@abc\"," +
                        "\"password\" : \"pass\"," +
                        "}"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.token", Matchers.notNullValue()))
                .andReturn();
        token = mapper.convertValue(result.getResponse().getContentAsString(), JsonObject.class).get("token").getAsString();

    }


    @Test
    public void save_success() throws Exception {
        mockMvc.perform(MockMvcRequestBuilders.get("/tasks")
                        .header("Authorization", "Bearer " + token)
                .content("{\"content\" : \"" + RandomString.make(100) + "\",}"))
                .andExpect(status().isOk());

    }

}
