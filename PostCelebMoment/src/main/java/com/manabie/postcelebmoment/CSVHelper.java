package com.manabie.postcelebmoment;

import java.io.BufferedReader;
import java.io.ByteArrayOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileWriter;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.nio.file.Files;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Iterator;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

import org.apache.commons.csv.CSVFormat;
import org.apache.commons.csv.CSVParser;
import org.apache.commons.csv.CSVPrinter;
import org.apache.commons.csv.CSVRecord;
import org.apache.commons.csv.QuoteMode;
import org.springframework.core.io.ClassPathResource;
import org.springframework.web.multipart.MultipartFile;

public class CSVHelper {
	public static String TYPE = "text/csv";
	static String[] HEADERs = { "userId", "lastPostDate", "limitation","counter" };

	public static boolean hasCSVFormat(MultipartFile file) {
		if (!TYPE.equals(file.getContentType())) {
			return false;
		}
		return true;
	}

	public static LinkedHashMap<String, UserTracking> csvToUserTracking() {
		File f = null;
		try {
			f = new ClassPathResource("user-tracking.csv").getFile();
		} catch (IOException e2) {
			e2.printStackTrace();
		}
		InputStream is = null;
		try {
			is = new FileInputStream(f);
		} catch (FileNotFoundException e1) {
			e1.printStackTrace();
		}
		try (BufferedReader fileReader = new BufferedReader(new InputStreamReader(is, "UTF-8"));
				CSVParser csvParser = new CSVParser(fileReader,
						CSVFormat.DEFAULT.withFirstRecordAsHeader().withIgnoreHeaderCase().withTrim());) {
			LinkedHashMap<String, UserTracking> userTracks = new LinkedHashMap<String, UserTracking>();
			Iterable<CSVRecord> csvRecords = csvParser.getRecords();
			for (CSVRecord csvRecord : csvRecords) {
				UserTracking userTrack = new UserTracking();
				userTrack.setUserId(csvRecord.get(0));
				userTrack.setLastPostDate(LocalDateTime.parse(csvRecord.get(1), DateTimeFormatter.ofPattern("dd-MMM-yyyy HH:mm")));
				userTrack.setLimitation(Integer.valueOf(csvRecord.get(2)));
				userTrack.setCounter(Integer.valueOf(csvRecord.get(3)));
				userTracks.put(csvRecord.get(0), userTrack);
			}
			return userTracks;
		} catch (IOException e) {
			throw new RuntimeException("fail to parse CSV file: " + e.getMessage());
		}
	}
	
	
	public static void userTrackingsToCSV(LinkedHashMap<String, UserTracking> userTrackings) {
		final CSVFormat format = CSVFormat.DEFAULT.withQuoteMode(QuoteMode.MINIMAL);
		DateTimeFormatter formatter = DateTimeFormatter.ofPattern("dd-MMM-yyyy HH:mm");
		try (ByteArrayOutputStream out = new ByteArrayOutputStream();
				CSVPrinter csvPrinter = new CSVPrinter(new PrintWriter(out), format);) {
			List<String> header = Arrays.asList(HEADERs);
			csvPrinter.printRecord(header);
			for (Map.Entry<String, UserTracking> entry : userTrackings.entrySet()) {
				List<String> data = Arrays.asList(entry.getKey(), entry.getValue().getLastPostDate().format(formatter),
						String.valueOf(entry.getValue().getLimitation()),
						String.valueOf(entry.getValue().getCounter()));
				csvPrinter.printRecord(data);
			}
			csvPrinter.flush();
			try {
				File f = new File("D:\\com-manabie-exam\\PostCelebMoment\\src\\main\\resources\\user-tracking.csv");
				Files.write(f.toPath(), out.toByteArray());
			} catch (IOException e2) {
				e2.printStackTrace();
			}
		} catch (IOException e) {
			e.printStackTrace();
		}
	}
	
}