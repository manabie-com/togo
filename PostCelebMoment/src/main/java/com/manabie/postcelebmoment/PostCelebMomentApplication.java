package com.manabie.postcelebmoment;

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Date;
import java.util.LinkedHashMap;

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
	private LinkedHashMap<String, UserTracking> userTrackList = new LinkedHashMap<String, UserTracking>();
	
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
		this.userTrackList = CSVHelper.csvToUserTracking();
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

	private boolean isExceedLimit(String userId) {
		DateTimeFormatter dtf = DateTimeFormatter.ofPattern("dd-MMM-yyyy HH:mm");  
		LocalDateTime today = LocalDateTime.now();
		//insert new user incase user not existed
		if(!this.userTrackList.containsKey(userId)) {
			UserTracking newUser = new UserTracking(userId, today, 5, 1);
			this.userTrackList.put(userId, newUser);
			return false;
		}else {
			//compare lastPostDate and counter
			LocalDateTime lastPostDate = this.userTrackList.get(userId).getLastPostDate();
			if(lastPostDate.getDayOfYear() < today.getDayOfYear() || lastPostDate.getYear() < today.getYear()) {
				//reset counter for first post on day
				this.userTrackList.get(userId).setCounter(1);
				this.userTrackList.get(userId).setLastPostDate(today);
				return false;
			}else {
				int counter = this.userTrackList.get(userId).getCounter();
				if(counter + 1 <= this.userTrackList.get(userId).getLimitation()) {
					//counter + 1
					this.userTrackList.get(userId).setCounter(counter + 1);
					return false;
				}else {
					return true;
				}
			}
		}
	}
	
	public String uploadFile(MultipartFile multipartFile, String userId) {

	    String fileUrl = "";
	    try {
	    	//TODO validate userId
	    	if(isExceedLimit(userId)) {
	    		return "You only can post " + this.userTrackList.get(userId).getLimitation() + " photos per day";
	    	}
	        File file = convertMultiPartToFile(multipartFile);
	        String fileName = generateFileName(multipartFile);
	        fileUrl = endpointUrl + "/" + bucketName + "/" + fileName;
	        uploadFileTos3bucket(fileName, file);
	        file.delete();
	        CSVHelper.userTrackingsToCSV(this.userTrackList);
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
