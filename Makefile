day = $(shell date +'%-d')

new-day:
	@echo "Creating new file structure for day" $(day)"..."

	@if [ $(day) -lt 10 ] ; then \
  		mkdir calendar/day-0$(day); \
  		cp static-template/dayxx.go calendar/day-0$(day)/day0$(day).go; \
  		cp static-template/dayxx_pt02.go calendar/day-0$(day)/day0$(day)_pt02.go; \
  	else \
  		mkdir calendar/day-$(day); \
		cp static-template/dayxx.go calendar/day-$(day)/day$(day).go; \
		cp static-template/dayxx_pt02.go calendar/day-$(day)/day$(day)_pt02.go; \
    fi
	@echo "Files successfully created.. happy hacking :)"
	@echo "INFO: puzzle input still needs to be fetched"
	@git add calendar/


