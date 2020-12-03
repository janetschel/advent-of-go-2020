day = $(shell date +'%d')
padded_day=$(shell printf '%02d' $(day))

new:
	@echo "Creating new file structure for day" $(day)"..."

	@mkdir calendar/day-$(padded_day); \
	cp template calendar/day-$(padded_day)/day$(padded_day).go; \
	cp template calendar/day-$(padded_day)/day$(padded_day)_pt02.go; \
	touch calendar/day-$(padded_day)/README.md;

	@echo "Files successfully created.. happy hacking :)"
	@echo "INFO: puzzle input still needs to be fetched"
	@git add calendar/
