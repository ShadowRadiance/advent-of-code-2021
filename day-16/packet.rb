class Packet
  attr_reader :version, :type
  def initialize(version, type)
    @version = version
    @type = type
    @subpackets = []
  end

  def version_sum
    version + subpackets.sum(&:version_sum)
  end

  protected

  attr_reader :subpackets

  def slurp(bitstream, count)
    bitstream.pop(count)
  end
end


