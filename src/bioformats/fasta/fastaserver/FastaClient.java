import java.net.Socket;
import java.util.List;
import java.util.ArrayList;
import java.io.InputStream;
import java.io.OutputStream;

/**
 * Communicates with a fasta server, for retrieval of genetic sequences.
 */
class FastaClient {
	/** Port on which the server listens. */
	private int port;
	
	/** Creates a client on the server's default port 1912. */
	public Client() {
		this(1912);
	}
	
	/** Creates a client on the given port. */
	public Client(int port) {
		this.port = port;
	}
	
	/**
	 * Returns the subsequence for the given parameters. Throws an exception
	 * if the chromosome name doesn't exist or the position is out of bounds.
	 * @param chr Chromosome name. Must match the name in the fasta file.
	 * @param start 0-based start position.
	 * @param length Length of the subsequence to fetch.
	 * @return The required subsequence, or null if an unexpected error
	 * occurred.
	 */
	public String getSequence(String chr, int start, int length) {
		try (
			Socket socket = new Socket("localhost", port)
		) {
			OutputStream out = socket.getOutputStream();
			InputStream in = socket.getInputStream();
			
			// Send message to server.
			out.write((chr + "," + start + "," + length + ";").getBytes());
			
			// Read response.
			List<Byte> responseList = new ArrayList<Byte>();
			for (int b = in.read(); b != -1; b = in.read()) {
				responseList.add((byte)b);
			}
			
			// Convert to string.
			byte[] bytes = new byte[responseList.size()];
			for (int i = 0; i < bytes.length; i++) {
				bytes[i] = responseList.get(i);
			}
			
			String response = new String(bytes);
			
			// 1  signifies error.
			if (response.charAt(0) == '1') {
				throw new FastaClientException(response.substring(1));
			}
			
			return response.substring(1);
		} catch (FastaClientException e) {
			throw e;
		} catch (Exception e) {
			return null;
		}
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
}
