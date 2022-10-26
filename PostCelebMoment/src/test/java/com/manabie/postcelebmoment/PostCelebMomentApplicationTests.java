package com.manabie.postcelebmoment;

import java.io.File;
import java.io.IOException;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.context.ApplicationContext;
import org.springframework.core.io.ClassPathResource;

import com.amazonaws.services.s3.waiters.AmazonS3Waiters;

@SpringBootTest
class PostCelebMomentApplicationTests {
	
	@Autowired
	private ApplicationContext context;
	
	@Test
	void contextLoads() {
		Assertions.assertTrue(context.getBean(PostCelebMomentApplication.class) != null);
	}
	
	@Test
	void awsConfigLoad() {
		Assertions.assertEquals("AKIA3L76SIDYZYCYKG43", context.getBean(PostCelebMomentApplication.class).getAccessKey());
		Assertions.assertEquals("manabie-landing-celebmoment", context.getBean(PostCelebMomentApplication.class).getBucketName());
		Assertions.assertEquals("pTr7+2FntL+DDXMOpxo4QrbJryMBLO4C7BG4BKaV", context.getBean(PostCelebMomentApplication.class).getSecretKey());
	}
	
	@Test
	void awsTestConnection() {
		Assertions.assertTrue(context.getBean(PostCelebMomentApplication.class).getS3client().listBuckets().size() > 0);
	}
	
	@Test
	void uploadFileTos3bucket() {
		 try {
			File f = new ClassPathResource("messi.jpg").getFile();
			Assertions.assertTrue(context.getBean(PostCelebMomentApplication.class).uploadFileTos3bucket("messi.jpg", f) == true);
		} catch (IOException e) {
			e.printStackTrace();
		}
	}
	
}
