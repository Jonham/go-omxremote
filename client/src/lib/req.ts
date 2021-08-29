export function getHostName(): string {
	if (location.port === '3000') {
		return 'http://localhost:31415';
	}
	return '';
}
