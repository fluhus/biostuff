import java.net.Socket;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.util.HashMap;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;

import org.json.simple.*;
import org.json.simple.parser.*;

/**
 * Communicates with a fasta server, for retrieval of genetic sequences.
 */
public class FastaClient {
	/**
	 * Returns the subsequence for the given parameters.
	 * @param port Port number of the listening server.
	 * @param chr Chromosome name. Must match the name in the fasta file.
	 * @param start 0-based start position.
	 * @param length Length of the subsequence to fetch.
	 * @return The required subsequence, or null if an unexpected error
	 * occurred.
	 */
	public static String getSequence(int port, String chr, int start,
			int length) {
		try (Socket socket = new Socket("localhost", port)) {
			OutputStream out = socket.getOutputStream();
			InputStream in = socket.getInputStream();
			
			// Send request.
			out.write(String.format("{\"type\":\"sequence\",\"sequence\":" +
					"{\"chr\":\"%s\",\"start\":%d,\"length\":%d}}",
					chr, start, length).getBytes());
			
			// Read response.
			JSONObject res = (JSONObject)JSONValue
					.parse(new InputStreamReader(in));
			
			// Check for error.
			if (res.containsKey("error")) {
				throw new FastaClientException((String)res.get("error"));
			}
			
			return (String) res.get("sequence");
		} catch (Exception e) {
			return null;
		}
	}
	
	/**
	 * Returns the subsequence for the given parameters. Using the server's
	 * default port.
	 * @param chr Chromosome name. Must match the name in the fasta file.
	 * @param start 0-based start position.
	 * @param length Length of the subsequence to fetch.
	 * @return The required subsequence, or null if an unexpected error
	 * occurred.
	 */
	public static String getSequence(String chr, int start, int length) {
		return getSequence(1912, chr, start, length);
	}
	
	/**
	 * Returns a map of available chromosomes and their lengths.
	 * @param port Port number of the listening server.
	 * @return A map where each key is a chromosome name and each value is
	 * the corresponding length.
	 */
	public static Map<String, Long> getChromosomes(int port) {
		try (Socket socket = new Socket("localhost", port)) {
			OutputStream out = socket.getOutputStream();
			InputStream in = socket.getInputStream();
			
			// Send request.
			out.write("{\"type\":\"meta\"}".getBytes());
			
			// Read response.
			JSONObject res = (JSONObject)JSONValue
					.parse(new InputStreamReader(in));
			
			// Check for error.
			if (res.containsKey("error")) {
				throw new FastaClientException((String)res.get("error"));
			}
			
			// Convert to map.
			Map<String, Long> result = new HashMap<>();
			for (Object key : res.keySet()) {
				result.put((String)key, (Long)res.get(key));
			}
			
			return result;
		} catch (Exception e) {
			return null;
		}
	}
	
	/**
	 * Returns a map of available chromosomes and their lengths. Contacts the
	 * default port.
	 * @return A map where each key is a chromosome name and each value is
	 * the corresponding length.
	 */
	public static Map<String, Long> getChromosomes() {
		return getChromosomes(1912);
	}
	
	/** Indicates errors returned by the fasta server. */
	public static class FastaClientException extends RuntimeException {
		public FastaClientException() {
			super();
		}
		public FastaClientException(String message) {
			super(message);
		}
	}
	
	public static void main(String[] args) {
		System.out.println("s=" + getSequence("chrI", 100, 10));
		System.out.println(getChromosomes(1912));
	}
}
