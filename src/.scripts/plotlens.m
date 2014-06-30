clear;
a = load('a.txt');

y = a(:,1);
x = a(:,3:end);

positives = x(y == 1,:);
negatives = x(y == -1,:);

plot(negatives', '^b');
plot(positives(1:100,:)', 'pr')
axis([-1 10 -1 10]);
