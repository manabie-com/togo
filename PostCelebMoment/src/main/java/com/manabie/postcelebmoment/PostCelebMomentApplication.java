package com.manabie.postcelebmoment;

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Date;

import javax.annotation.PostConstruct;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.multipart.MultipartFile;

import com.amazonaws.AmazonServiceException;
import com.amazonaws.SdkClientException;
import com.amazonaws.auth.AWSStaticCredentialsProvider;
import com.amazonaws.auth.BasicAWSCredentials;
import com.amazonaws.services.s3.AmazonS3;
import com.amazonaws.services.s3.AmazonS3ClientBuilder;
import com.amazonaws.services.s3.model.CannedAccessControlList;
import com.amazonaws.services.s3.model.PutObjectRequest;

@SpringBootApplication
public class PostCelebMomentApplication {
	
	private AmazonS3 s3client;

	@Value("${amazonProperties.endpointUrl}")
	private String endpointUrl;
	@Value("${amazonProperties.bucketName}")
	private String bucketName;
	@Value("${amazonProperties.accessKey}")
	private String accessKey;
	@Value("${amazonProperties.secretKey}")
	private String secretKey;
	@Value("${amazonProperties.region}")
	private String region;

	@PostConstruct
	private void initializeAmazon() {
		BasicAWSCredentials creds = new BasicAWSCredentials(this.accessKey, this.secretKey); 
		this.s3client = AmazonS3ClientBuilder.standard().withCredentials(new AWSStaticCredentialsProvider(creds)).withRegion(this.region).build();
	}
	
	
	private File convertMultiPartToFile(MultipartFile file) throws IOException {
	    File convFile = new File(file.getOriginalFilename());
	    FileOutputStream fos = new FileOutputStream(convFile);
	    fos.write(file.getBytes());
	    fos.close();
	    return convFile;
	}
	
	private String generateFileName(MultipartFile multiPart) {
	    return new Date().getTime() + "-" + multiPart.getOriginalFilename().replace(" ", "_");
	}
	
	public boolean uploadFileTos3bucket(String fileName, File file) {
	    try {
	    	s3client.putObject(new PutObjectRequest(bucketName, fileName, file)
		            .withCannedAcl(CannedAccessControlList.PublicRead));
	        return true;
	      } catch (AmazonServiceException e) {
	        return false;
	      } catch (SdkClientException e) {
	        return false;
	      }
	}
	
	public String uploadFile(MultipartFile multipartFile) {

	    String fileUrl = "";
	    try {
	        File file = convertMultiPartToFile(multipartFile);
	        String fileName = generateFileName(multipartFile);
	        fileUrl = endpointUrl + "/" + bucketName + "/" + fileName;
	        uploadFileTos3bucket(fileName, file);
	        file.delete();
	    } catch (Exception e) {
	       e.printStackTrace();
	       return "";
	    }
	    return fileUrl;
	}
	
	public String getEndpointUrl() {
		return endpointUrl;
	}


	public void setEndpointUrl(String endpointUrl) {
		this.endpointUrl = endpointUrl;
	}


	public String getBucketName() {
		return bucketName;
	}


	public void setBucketName(String bucketName) {
		this.bucketName = bucketName;
	}


	public String getAccessKey() {
		return accessKey;
	}


	public void setAccessKey(String accessKey) {
		this.accessKey = accessKey;
	}


	public String getSecretKey() {
		return secretKey;
	}


	public void setSecretKey(String secretKey) {
		this.secretKey = secretKey;
	}


	public String getRegion() {
		return region;
	}


	public void setRegion(String region) {
		this.region = region;
	}

	public AmazonS3 getS3client() {
		return s3client;
	}

	public void setS3client(AmazonS3 s3client) {
		this.s3client = s3client;
	}

	public static void main(String[] args) {
		SpringApplication.run(PostCelebMomentApplication.class, args);
	}
	
}
