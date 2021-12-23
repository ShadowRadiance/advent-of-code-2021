# frozen_string_literal: true

require "set"

class Recombobulator
  # assume scanner 0 is "correct"
  # locate other scanners and turn them to the correct orientation
  # find all the beacons

  attr_reader :beacons, :scanners

  # array<scanner> -> nil
  def initialize(scanners)
    @beacons = Set.new
    @scanners = scanners

    @base_scanner = scanners.first
    merge_scanner(@base_scanner)
    @remaining = scanners.slice(1..)
  end

  # nil -> nil
  def recombobulate
    cycled_to_back = 0
    while @remaining.size.positive?
      scanner = @remaining.shift
      if match_scanner(scanner)
        merge_scanner(scanner)
      else
        @remaining.push(scanner)
        cycled_to_back += 1
        raise "Panic" if cycled_to_back > 1_000_000
      end
    end
  end

  private

  # scanner -> true/false
  def match_scanner(scanner)
    (0..23).each do |index|
      # puts "CHECKING SCANNER #{scanner.identifier} orientation #{index}"
      beacons = scanner.beacons_in_orientation(index)
      transform = match_beacons(beacons)
      if transform
        # puts "FOUND A MATCH with orientation #{index} - SCANNER IS AT #{transform}"
        scanner.location = transform
        scanner.orientation = index
        return true
      end
    end
    false
  end

  # return Vector from main scanner set or nil if we can't figure out a match
  #
  # array<array<int>> -> vector
  def match_beacons(other_beacons)
    other_beacons.each.with_index do |main_other_beacon, main_other_beacon_index|
      remaining_other_beacons = other_beacons[(main_other_beacon_index+1)..]

      beacons.each do |main_beacon|
        # figure out the vector (transform) to shift main_other_beacon to main_beacon
        transform = main_beacon - main_other_beacon

        # apply the transform to each of the other beacons
        #  - if 11+ of them match beacons in self, the set is a match
        matched = 0
        remaining_other_beacons.each do |other_beacon|
          transformed_other = other_beacon + transform
          if beacons.include?(transformed_other)
            # this is a match
            matched += 1
            # return the transform vector if we have enough
            return transform if matched >= 11
          end
        end
      end
    end

    nil
  end

  # scanner -> void
  def merge_scanner(scanner)
    scanner.beacons.each do |beacon|
      @beacons << beacon + scanner.location
    end
  end
end
