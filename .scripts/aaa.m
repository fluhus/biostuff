clear; close all;
a = importdata('dist_sim.txt');

size(a, 1)

for ii = 1:size(a, 1)
%    figure('name', sprintf('chromosome %d', ii));
%    plot(a(ii,2:1+a(ii,1)),'o');
max(a(ii,2:1+a(ii,1)))
end
