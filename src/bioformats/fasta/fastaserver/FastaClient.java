import java.io.*;
import java.net.URL;
import java.util.*;
import java.util.regex.*;

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
		try {
			// Connect to server.
			String address = String.format(
					"http://localhost:%d/sequence?chr=%s&start=%d&length=%d",
					port, chr, start, length);
			InputStream in = (new URL(address)).openStream();
			
			// Parse response.
			StringBuilder sb = new StringBuilder();
			for (int i = in.read(); i != -1; i = in.read()) {
				sb.append((char)i);
			}
			String result = sb.toString();
			in.close();
			
			// Handle error.
			if (result.startsWith("Error: ")) {
				throw new FastaClientException(result.substring(7));
			}
			
			return result;
		} catch (IOException e) {
			e.printStackTrace();
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
		try {
			// Connect to server.
			String address = String.format("http://localhost:%d/meta", port);
			InputStream in = (new URL(address)).openStream();
			Scanner scanner = new Scanner(in);
			
			// Parse response.
			Pattern pat = Pattern.compile("^(.*): (\\d+)$");
			Map<String, Long> result = new HashMap<>();
			while (scanner.hasNext()) {
				String line = scanner.nextLine();
				Matcher m = pat.matcher(line);
				m.find();
				
				long length = Long.valueOf(m.group(2));
				result.put(m.group(1), length);
			}
			
			in.close();
			
			return result;
		} catch (Exception e) {
			e.printStackTrace();
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
		System.out.println("s=" + getSequence("amit", 10, 10));
		System.out.println(getChromosomes(1912));
	}
}
