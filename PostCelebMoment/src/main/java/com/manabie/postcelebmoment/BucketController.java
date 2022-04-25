package com.manabie.postcelebmoment;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestPart;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

@RestController
@RequestMapping("/api/storage/")
public class BucketController {

    private PostCelebMomentApplication celebMomentClient;

    @Autowired
    BucketController(PostCelebMomentApplication client) {
        this.celebMomentClient = client;
    }

    @PostMapping("/celebMoment")
    public String uploadFile(@RequestPart(value = "file") MultipartFile file, @RequestPart(value = "userId") String userId) {
        return this.celebMomentClient.uploadFile(file, userId);
    }
}